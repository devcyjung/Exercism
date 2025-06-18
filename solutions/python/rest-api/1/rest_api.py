from json import dumps, loads
from typing import Literal

GET_URL_t = Literal["/users"]
POST_URL_t = Literal["/add", "/iou"]

class RestAPI:
    def __init__(self, database: dict = {}) -> None:
        self._db = {user.get('name'): user \
                    for user in database.get('users', [])}

    def get(self, url: GET_URL_t, payload: str | None = None) -> str:
        match url:
            case "/users":
                return self._get_users(payload) \
                    if payload else self._get_all_users()
            case _:
                raise ValueError(f'Invalid GET URL: {url}')

    def post(self, url: POST_URL_t, payload: str = "{}") -> str:
        match url:
            case "/add":
                return self._add_user(payload)
            case "/iou":
                return self._update_iou(payload)
            case _:
                raise ValueError(f'Invalid POST URL: {url}')

    def _get_all_users(self) -> str:
        return dumps({'users': [*self._db.values()]})

    def _get_users_from_list(self, names: list) -> str:
        return dumps({'users': [self._db.get(name) for name in names if name in self._db]})
    
    def _get_users(self, payload: str) -> str:
        return self._get_users_from_list(loads(payload).get('users', []))

    def _add_user(self, payload: str) -> str:
        name: str = loads(payload).get('user', '')
        user: dict = RestAPI._default_user(name)
        return dumps(self._db.setdefault(name, user))

    def _update_iou(self, payload: str) -> str:
        pl: dict = loads(payload)
        lender_name: str = pl.get('lender', '')
        borrower_name: str = pl.get('borrower', '')
        amount: float = pl.get('amount', 0.0)
        lender: dict = self._db.get(lender_name, RestAPI._default_user(lender_name))
        borrower: dict = self._db.get(borrower_name, RestAPI._default_user(borrower_name))
        lender['balance'] += amount
        borrower['balance'] -= amount
        lend: float = lender['owed_by'].pop(borrower_name, 0.0) - lender['owes'].pop(borrower_name, 0.0)
        borrower['owed_by'].pop(lender_name, 0.0)
        borrower['owes'].pop(lender_name, 0.0)
        lend += amount
        if lend > 0:
            lender['owed_by'][borrower_name] = lend
            borrower['owes'][lender_name] = lend
        if lend < 0:
            borrower['owed_by'][lender_name] = -lend
            lender['owes'][borrower_name] = -lend
        self._db |= {lender_name: lender, borrower_name: borrower}
        return self._get_users_from_list(sorted([lender_name, borrower_name]))
    
    @staticmethod
    def _default_user(name: str) -> dict:
        return {'name': name, 'owes': {}, 'owed_by': {}, 'balance': 0.0}
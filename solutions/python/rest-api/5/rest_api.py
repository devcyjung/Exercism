from json import dumps, loads
from typing import Literal

GET_URL_t = Literal["/users"]
POST_URL_t = Literal["/add", "/iou"]

class RestAPI:
    def __init__(self, database: dict | None = None) -> None:
        self._db = {user.get('name'): user \
                    for user in (database or {}).get('users', [])}

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
        return dumps({'users': [self._db.get(name) for name in names]})
    
    def _get_users(self, payload: str) -> str:
        return self._get_users_from_list(loads(payload).get('users', []))

    def _add_user(self, payload: str) -> str:
        name: str = loads(payload).get('user', '')
        user: dict = RestAPI._default_user(name)
        return dumps(self._db.setdefault(name, user))

    def _update_iou(self, payload: str) -> str:
        pl: dict = loads(payload)
        l_name: str = pl.get('lender', '')
        b_name: str = pl.get('borrower', '')
        amount: float = pl.get('amount', 0.0)
        lender: dict = self._db.get(l_name, RestAPI._default_user(l_name))
        borrower: dict = self._db.get(b_name, RestAPI._default_user(b_name))
        lender['balance'] += amount
        borrower['balance'] -= amount
        lend: float = lender['owed_by'].pop(b_name, 0.0) \
            - lender['owes'].pop(b_name, 0.0)
        borrower['owed_by'].pop(l_name, 0.0)
        borrower['owes'].pop(l_name, 0.0)
        lend += amount
        if lend > 0:
            lender['owed_by'][b_name] = lend
            borrower['owes'][l_name] = lend
        if lend < 0:
            borrower['owed_by'][l_name] = -lend
            lender['owes'][b_name] = -lend
        self._db |= {l_name: lender, b_name: borrower}
        return self._get_users_from_list(sorted([l_name, b_name]))
    
    @staticmethod
    def _default_user(name: str) -> dict:
        return {'name': name, 'owes': {}, 'owed_by': {}, 'balance': 0.0}
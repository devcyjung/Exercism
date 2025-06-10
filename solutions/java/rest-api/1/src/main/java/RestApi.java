import java.util.HashMap;
import java.util.Map;
import org.json.JSONArray;
import org.json.JSONObject;

class RestApi {
    private final Map<String, Map<String, Double>> netBorrow = new HashMap<>();

    RestApi(User... users) {
        for (var user: users) {
            netBorrow.putIfAbsent(user.name(), new HashMap<>());
            var userEntry = netBorrow.get(user.name());
            for (var iou: user.owes()) {
                userEntry.put(iou.name, userEntry.getOrDefault(iou.name, 0.0) + iou.amount);
            }
            for (var iou: user.owedBy()) {
                userEntry.put(iou.name, userEntry.getOrDefault(iou.name, 0.0) - iou.amount);
            }
        }
        setOff();
    }

    private void setOff() {
        netBorrow.entrySet().forEach(netBorrowEntry -> {
            netBorrowEntry.getValue().entrySet().removeIf(e -> e.getValue() == 0.0);
        });
    }

    String get(String url) {
        return switch (url) {
            case "/users" -> getAllUsers().toString();
            case "/add" -> throw new UnsupportedOperationException("/add is not a GET endpoint");
            case "/iou" -> throw new UnsupportedOperationException("/iou is not a GET endpoint");
            default -> throw new IllegalArgumentException("Bad URL: " + url);
        };
    }

    String get(String url, JSONObject payload) {
        return switch (url) {
            case "/users" -> getUsers(payload.optJSONArray("users")).toString();
            case "/add" -> throw new UnsupportedOperationException("/add is not a GET endpoint");
            case "/iou" -> throw new UnsupportedOperationException("/iou is not a GET endpoint");
            default -> throw new IllegalArgumentException("Bad URL: " + url);
        };
    }

    String post(String url, JSONObject payload) {
        return switch (url) {
            case "/users" -> throw new UnsupportedOperationException("/users is not a POST endpoint");
            case "/add" -> addUser(payload.optString("user")).toString();
            case "/iou" -> addBorrow(
                payload.optString("borrower"), payload.optString("lender"), payload.optDouble("amount")
            ).toString();
            default -> throw new IllegalArgumentException("Bad URL: " + url);
        };
    }

    private JSONObject singleUser(String borrower, Map<String, Double> lenders) {
        var user = new JSONObject().put("name", borrower);
        var owes = new JSONObject();
        var owedBy = new JSONObject();
        double[] balance = {0.0};
        lenders.forEach((lender, amount) -> {
            balance[0] -= amount;
            if (amount < 0.0) {
                owedBy.put(lender, -amount);
            }
            if (amount > 0.0) {
                owes.put(lender, amount);
            }
        });
        return user.put("owes", owes).put("owedBy", owedBy).put("balance", balance[0]);
    }

    JSONObject getAllUsers() {
        var users = new JSONArray();
        netBorrow.forEach((borrower, lenders) -> {
            users.put(singleUser(borrower, lenders));
        });
        return new JSONObject().put("users", users);
    }

    JSONObject getUsers(JSONArray usernames) {
        var users = new JSONArray();
        usernames.forEach(borrowerObj -> {
            String borrower = (String) borrowerObj;
            var lenders = netBorrow.get(borrower);
            if (lenders != null) {
                users.put(singleUser(borrower, lenders));
            }
        });
        return new JSONObject().put("users", users);
    }

    JSONObject addUser(String username) {
        netBorrow.putIfAbsent(username, new HashMap<>());
        return singleUser(username, netBorrow.get(username));
    }

    JSONObject addBorrow(String borrower, String lender, double amount) {
        netBorrow.putIfAbsent(borrower, new HashMap<>());
        netBorrow.putIfAbsent(lender, new HashMap<>());
        var borrowerData = netBorrow.get(borrower);
        var lenderData = netBorrow.get(lender);
        borrowerData.put(lender, borrowerData.getOrDefault(lender, 0.0) + amount);
        lenderData.put(borrower, lenderData.getOrDefault(borrower, 0.0) - amount);
        setOff();
        if (borrower.compareTo(lender) < 0) {
            return getUsers(new JSONArray().put(borrower).put(lender));
        } else {
            return getUsers(new JSONArray().put(lender).put(borrower));
        }
    }
}
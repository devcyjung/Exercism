import java.util.Map;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;

public class DialingCodes {
    private final Map<Integer, String> codes = new ConcurrentHashMap();
    private final Set<Map.Entry<Integer, String>> entries = codes.entrySet();

    public Map<Integer, String> getCodes() {
        return Map.copyOf(codes);
    }

    public void setDialingCode(Integer code, String country) {
        codes.put(code, country);
    }

    public String getCountry(Integer code) {
        return codes.get(code);
    }

    public void addNewDialingCode(Integer code, String country) {
        if (!codes.containsValue(country)) {
            codes.putIfAbsent(code, country);
        }
    }

    public Integer findDialingCode(String country) {
        for (var entry: entries) {
            if (entry.getValue().equals(country)) {
                return entry.getKey();
            }
        }
        return null;
    }

    public void updateCountryDialingCode(Integer code, String country) {
        var oldCode = findDialingCode(country);
        if (oldCode != null) {
            codes.remove(oldCode);
            codes.put(code, country);
        }
    }
}
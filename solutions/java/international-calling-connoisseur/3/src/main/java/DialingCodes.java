import java.util.Collection;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class DialingCodes {
    private final Map<Integer, String> codes = new ConcurrentHashMap();
    private final Lock lock = new ReentrantLock(true);

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
        for (var entry: codes.entrySet()) {
            if (entry.getValue().equals(country)) {
                return entry.getKey();
            }
        }
        return null;
    }

    public void updateCountryDialingCode(Integer code, String country) {
        lock.lock();
        try {
            if (codes.values().removeIf(country::equals)) {
                codes.put(code, country);
            }   
        } finally {
            lock.unlock();
        }
    }
}
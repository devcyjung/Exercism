import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public record User(String name, List<Iou> owes, List<Iou> owedBy) {

    public User {
        owes = Collections.unmodifiableList(new ArrayList<>(owes));
        owedBy = Collections.unmodifiableList(new ArrayList<>(owedBy));
    }

    public static Builder builder() {
        return new Builder();
    }

    public static class Builder {
        private String name;
        private final List<Iou> owes = new ArrayList<>();
        private final List<Iou> owedBy = new ArrayList<>();

        public Builder setName(String name) {
            this.name = name;
            return this;
        }

        public Builder owes(String name, double amount) {
            owes.add(new Iou(name, amount));
            return this;
        }

        public Builder owedBy(String name, double amount) {
            owedBy.add(new Iou(name, amount));
            return this;
        }

        public User build() {
            return new User(name, owes, owedBy);
        }
    }
}
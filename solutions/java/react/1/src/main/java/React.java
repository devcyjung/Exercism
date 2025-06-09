import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Collections;
import java.util.IdentityHashMap;
import java.util.List;
import java.util.Set;
import java.util.function.Consumer;
import java.util.function.Function;
import java.util.stream.Collectors;

public class React {

    public static class Cell<T> {
        protected T value;
        protected Set<ComputeCell<T>> dependants = Collections.newSetFromMap(new IdentityHashMap<>());
        
        public T getValue() {
            return value;
        }

        protected void addDependant(ComputeCell<T> cell) {
            dependants.add(cell);
        }
    }

    public static class InputCell<T> extends Cell<T> {
        private InputCell(T value) {
            this.value = value;
        }
        
        public void setValue(T newValue) {
            if (value.equals(newValue)) {
                return;
            }
            value = newValue;
            var queue = new ArrayDeque<ComputeCell<T>>();
            var invokeList = new ArrayList<ComputeCell<T>>();
            var visited = Collections.newSetFromMap(new IdentityHashMap<ComputeCell<T>, Boolean>());
            queue.addAll(dependants);
            visited.addAll(dependants);
            ComputeCell<T> currentCell;
            T prevValue;
            while ((currentCell = queue.poll()) != null) {
                prevValue = currentCell.getValue();
                currentCell.updateValue();
                if (!prevValue.equals(currentCell.getValue())) {
                    invokeList.add(currentCell);
                }
                currentCell.dependants.stream()
                    .filter(dep -> !visited.contains(dep))
                    .forEach(dep -> queue.add(dep));
            }
            invokeList.forEach(ComputeCell::invokeCallbacks);
        }
    }

    public static class ComputeCell<T> extends Cell<T> {
        private Function<List<T>, T> function;
        private List<Cell<T>> dependencies;
        private Set<Consumer<T>> callbacks = Collections.newSetFromMap(new IdentityHashMap<>());

        private ComputeCell(Function<List<T>, T> function, List<Cell<T>> dependencies) {
            this.function = function;
            this.dependencies = dependencies;
            this.value = computeValue();
        }

        private T computeValue() {
            return function.apply(dependencies.stream()
                                  .map(Cell::getValue)
                                  .collect(Collectors.toList()));
        }

        public void updateValue() {
            value = computeValue();
        }

        private void invokeCallbacks() {
            callbacks.stream().forEach(cb -> cb.accept(value));
        }
        
        public void addCallback(Consumer<T> callback) {
            callbacks.add(callback);
        }

        public void removeCallback(Consumer<T> callback) {
            callbacks.remove(callback);
        }
    }

    public static <T> InputCell<T> inputCell(T initialValue) {
        return new InputCell<T>(initialValue);
    }

    public static <T> ComputeCell<T> computeCell(Function<List<T>, T> function, List<Cell<T>> cells) {
        var newComputeCell = new ComputeCell<T>(function, cells);
        cells.stream().forEach(c -> c.addDependant(newComputeCell));
        return newComputeCell;
    }
}
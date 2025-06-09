import java.util.Collection;
import java.util.HashSet;
import java.util.Set;
import java.util.stream.Collectors;

class GottaSnatchEmAll {

    static <T> Set<T> newCollection(Collection<T> cards) {
        return Set.copyOf(cards);
    }

    static <T> boolean addCard(T card, Collection<? super T> collection) {
        return collection.add(card);
    }

    static boolean canTrade(Collection<?> myCollection, Collection<?> theirCollection) {
        return !myCollection.isEmpty()
            && !theirCollection.isEmpty()
            && !myCollection.containsAll(theirCollection)
            && !theirCollection.containsAll(myCollection);
    }

    static <T> Set<T> commonCards(Collection<? extends Collection<T>> collections) {
        var iter = collections.iterator();
        if (!iter.hasNext()) {
            return Set.of();
        }
        Set<T> initial = new HashSet<T>(iter.next());
        for (var coll: collections) {
            initial.retainAll(coll);
        }
        return initial;
    }

    static <T> Set<T> allCards(Collection<? extends Collection<T>> collections) {
        return collections.stream().flatMap(Collection::stream).collect(Collectors.toSet());
    }
}
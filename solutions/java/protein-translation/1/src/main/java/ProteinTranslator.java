import java.util.ArrayList;
import java.util.List;

class ProteinTranslator {

    static List<String> translate(String rnaSequence) {
        List<String> result = new ArrayList<>();
        int index = 0;
        while (index < rnaSequence.length()) {
            if (index + 3 > rnaSequence.length()) {
                throw new IllegalArgumentException("Invalid codon");
            }
            String amino = toAmino(rnaSequence.substring(index, index + 3));
            if (amino == "STOP") {
                break;
            }
            result.add(amino);
            index += 3;
        }
        return result;
    }

    static String toAmino(String codon) {
        return switch (codon) {
            case "AUG" -> "Methionine";
            case "UUU", "UUC" -> "Phenylalanine";
            case "UUA", "UUG" -> "Leucine";
            case "UCU", "UCC", "UCA", "UCG" -> "Serine";
            case "UAU", "UAC" -> "Tyrosine";
            case "UGU", "UGC" -> "Cysteine";
            case "UGG" -> "Tryptophan";
            case "UAA", "UAG", "UGA" -> "STOP";
            default -> throw new IllegalArgumentException("Invalid codon");
        };
    }
}

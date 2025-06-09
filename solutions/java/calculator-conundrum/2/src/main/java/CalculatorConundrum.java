class CalculatorConundrum {
    public static String calculate(int operand1, int operand2, String operation) {
        return switch (operation) {
            case null ->
                throw new IllegalArgumentException("Operation cannot be null");
            
            case "" ->
                throw new IllegalArgumentException("Operation cannot be empty");
            
            case "+" -> 
                Integer.valueOf(operand1) + " + "
                + Integer.valueOf(operand2) + " = "
                + Integer.valueOf(operand1 + operand2);
                
            case "*" ->
                Integer.valueOf(operand1) + " * "
                + Integer.valueOf(operand2) + " = "
                + Integer.valueOf(operand1 * operand2);

            case "/" -> {
                try {
                    yield Integer.valueOf(operand1) + " / "
                        + Integer.valueOf(operand2) + " = "
                        + Integer.valueOf(operand1 / operand2);
                } catch (ArithmeticException ex) {
                    throw new IllegalOperationException(
                        "Division by zero is not allowed", ex
                    );
                }
            }
                
            default ->
                throw new IllegalOperationException(
                    "Operation '" + operation + "' does not exist"
                );
        };
    }
}
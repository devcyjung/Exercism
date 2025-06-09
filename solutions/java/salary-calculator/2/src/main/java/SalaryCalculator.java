public class SalaryCalculator {
    public static double salaryMultiplier(int daysSkipped) {
        return (daysSkipped >= 5) ? 0.85 : 1.0;
    }

    public static int bonusMultiplier(int productsSold) {
        return (productsSold < 20) ? 10 : 13;
    }

    public static double bonusForProductsSold(int productsSold) {
        return productsSold * bonusMultiplier(productsSold);
    }

    public static double finalSalary(int daysSkipped, int productsSold) {
        var salary = BASE_SALARY * salaryMultiplier(daysSkipped)
            + bonusForProductsSold(productsSold);
        return (salary > SALARY_CAP) ? SALARY_CAP : salary;
    }

    private static final double BASE_SALARY = 1000.0;
    private static final double SALARY_CAP = 2000.0;
}
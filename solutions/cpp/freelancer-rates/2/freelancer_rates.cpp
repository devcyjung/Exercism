#include <cmath>

const int BILLABLE_DAYS = 22;
const int WORK_HOUR = 8;

double daily_rate(double hourly_rate) {
    return WORK_HOUR * hourly_rate;
}

double apply_discount(double before_discount, double discount) {
    return before_discount * (1 - discount / 100);
}

int monthly_rate(double hourly_rate, double discount) {
    return ceil(apply_discount(daily_rate(hourly_rate), discount) * BILLABLE_DAYS);
}

int days_in_budget(int budget, double hourly_rate, double discount) {
    return floor(budget / apply_discount(daily_rate(hourly_rate), discount));
}
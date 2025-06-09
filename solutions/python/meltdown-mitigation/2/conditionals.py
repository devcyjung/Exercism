"""Functions to prevent a nuclear meltdown."""

def is_criticality_balanced(temperature, neutrons_emitted):
    return temperature < 800 and neutrons_emitted > 500 and temperature * neutrons_emitted < 5e5

def reactor_efficiency(voltage, current, theoretical_max_power):
    generated_power = voltage * current
    percent = (generated_power / theoretical_max_power) * 100
    if percent < 30:
        return "black"
    if percent < 60:
        return "red"
    if percent < 80:
        return "orange"
    return "green"

def fail_safe(temperature, neutrons_produced_per_second, threshold):
    value = temperature * neutrons_produced_per_second
    if value < 0.9 * threshold:
        return "LOW"
    if value <= 1.1 * threshold:
        return "NORMAL"
    return "DANGER"
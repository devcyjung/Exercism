enum Vulnerability {
    VULNERABLE, INVULNERABLE
}

class Fighter {
    Vulnerability getVulnerability() {
        if (isVulnerable()) {
            return Vulnerability.VULNERABLE;
        }
        return Vulnerability.INVULNERABLE;
    }

    boolean isVulnerable() {
        return true;
    }

    int getDamagePoints(Fighter fighter) {
        return 1;
    }

    @Override
    public String toString() {
        return String.format("%s is a %s",
            getClass().getSuperclass().getSimpleName(),
            getClass().getSimpleName());
    }
}

class Warrior extends Fighter {
    @Override
    boolean isVulnerable() {
        return false;    
    }

    @Override
    int getDamagePoints(Fighter fighter) {
        return switch(fighter.getVulnerability()) {
            case VULNERABLE -> 10;
            case INVULNERABLE -> 6;
        };
    }

    @Override
    public String toString() {
        return super.toString();
    }
}

class Wizard extends Fighter {
    private boolean isVulnerable = true;
    
    void prepareSpell() {
        isVulnerable = false;    
    }
    
    @Override
    boolean isVulnerable() {
        return isVulnerable;   
    }

    @Override
    int getDamagePoints(Fighter fighter) {
        return switch(getVulnerability()) {
            case VULNERABLE -> 3;
            case INVULNERABLE -> 12;
        };
    }

    @Override
    public String toString() {
        return super.toString();
    }
}
class ProductionRemoteControlCar
    implements
        RemoteControlCar,
        Comparable<ProductionRemoteControlCar> {
    private int distance = 0;
    private int victories = 0;
    private static int speed = 10;
    
    @Override
    public void drive() {
        distance += speed;
    }

    @Override
    public int getDistanceTravelled() {
        return distance;
    }

    @Override
    public int compareTo(ProductionRemoteControlCar other) {
        if (this.equals(other)) {
            return 0;
        }
        return -Integer.valueOf(getNumberOfVictories())
            .compareTo(Integer.valueOf(other.getNumberOfVictories()));
    }

    @Override
    public boolean equals(Object other) {
        return other != null && other.getClass() == getClass()
            && ((ProductionRemoteControlCar) other).getDistanceTravelled() == getDistanceTravelled()
            && ((ProductionRemoteControlCar) other).getNumberOfVictories() == getNumberOfVictories();
    }

    public int getNumberOfVictories() {
        return victories;
    }

    public void setNumberOfVictories(int numberOfVictories) {
        if (numberOfVictories >= 0) {
            victories = numberOfVictories;
        }
    }
}
class SpaceAge:
    _EARTH_YEAR = 31_557_600
    
    def __init__(self, seconds):
        self.seconds = seconds

    def on_earth(self):
        return round(self.seconds / SpaceAge._EARTH_YEAR, 2)

    def on_mercury(self):
        return round(self.seconds / 0.2408467 / SpaceAge._EARTH_YEAR, 2)

    def on_venus(self):
        return round(self.seconds / 0.61519726 / SpaceAge._EARTH_YEAR, 2)

    def on_mars(self):
        return round(self.seconds / 1.8808158 / SpaceAge._EARTH_YEAR, 2)

    def on_jupiter(self):
        return round(self.seconds / 11.862615 / SpaceAge._EARTH_YEAR, 2)

    def on_saturn(self):
        return round(self.seconds / 29.447498 / SpaceAge._EARTH_YEAR, 2)

    def on_uranus(self):
        return round(self.seconds / 84.016846 / SpaceAge._EARTH_YEAR, 2)

    def on_neptune(self):
        return round(self.seconds / 164.79132 / SpaceAge._EARTH_YEAR, 2)
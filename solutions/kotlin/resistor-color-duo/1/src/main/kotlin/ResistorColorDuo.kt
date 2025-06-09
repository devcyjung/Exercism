object ResistorColorDuo {
    fun value(vararg colors: Color): Int {
        require(colors.size >= 2) {
            "Must provide at least two colors"
        }
        return colors.get(0).ordinal * 10 + colors.get(1).ordinal
    }
}
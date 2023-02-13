package global

/* java

public enum Level {

    _0(new Color(0, 88, 248), new Color(63, 191, 255)),
    _1(new Color(0, 171, 0), new Color(184, 248, 24)),
    _2(new Color(219, 0, 205), new Color(248, 120, 248)),
    _3(new Color(0, 88, 248), new Color(91, 219, 87)),
    _4(new Color(231, 0, 91), new Color(88, 248, 152)),
    _5(new Color(88, 248, 152), new Color(107, 136, 255)),
    _6(new Color(248, 56, 0), new Color(127, 127, 127)),
    _7(new Color(107, 71, 255), new Color(171, 0, 35)),
    _8(new Color(0, 88, 248), new Color(248, 56, 0)),
    _9(new Color(248, 56, 0), new Color(255, 163, 71));


    private final Color c1, c2;
    Level(Color c1, Color c2) {
        this.c1 = c1;
        this.c2 = c2;
    }

    public static Level fromInt(int level) {
        return Level.valueOf("_" + Math.min(level, 9));
    }

    public Color getColor(Piece piece) {
        return Arrays.asList(Piece.I, Piece.J, Piece.O, Piece.S, Piece.T).contains(piece) ? c1 : c2;
    }

}
*/

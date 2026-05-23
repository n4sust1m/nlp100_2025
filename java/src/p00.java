class P00 {
  public static void main (String[] args) {
    String s1 = "パトカー";
    String s2 = "タクシー";

    StringBuilder result = new StringBuilder("");
    for (int i = 0; i < s1.length(); i++) {
      result.append( s1.charAt(i) );
      result.append( s2.charAt(i) );
    }
    System.out.println(result.toString());
  }
}
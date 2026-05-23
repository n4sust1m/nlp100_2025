class P01 {
  public static void main(String[] args) {
    String s1 = "パタトクカシーー";
    StringBuilder result = new StringBuilder("");

    int i = 1;
    do {
      result.append(s1.charAt(i));

      i += 2;
    } while (i < s1.length());

    System.out.println(result.toString());
  }
}

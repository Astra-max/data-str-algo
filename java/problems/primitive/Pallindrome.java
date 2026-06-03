public class Pallindrome {
	public static void main(String[] arg) {
		Pallindrome pl = new Pallindrome();
		System.out.println("num: " + pl.isPallindrome(1541));
	}
	public boolean isPallindrome(int num) {
		return Pallindrome.reverse(num) == num;
	}

	public static int reverse(int num) {
		int absVal = Math.abs(num);
		int results = 0;

		while (absVal > 0) {
			int lastVal = absVal % 10;
			results = results * 10 + lastVal;
			absVal /= 10;
		}
		return num > 0 ? results : -results;
	}
}

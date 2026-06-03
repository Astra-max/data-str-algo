public class ReverseNum {
	public static void main(String[] args) {
		ReverseNum rm = new ReverseNum();
		System.out.printf("reversed:: %d\n", rm.reverse(123));
	}
	public int reverse(int num) {
		int positiveVal = Math.abs(num);
		int results = 0;

		while (positiveVal > 0) {
			int lastVal = positiveVal % 10;
			results = results * 10 + lastVal;
			if (results >= Integer.MAX_VALUE) return 0;
			positiveVal /= 10;

		}
		return num < 0 ? -results : results;
	}
}

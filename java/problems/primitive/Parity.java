public class Parity {
	public static void main(String[] args) {
		Parity pr = new Parity();
		System.out.printf("parity: %d\n", pr.parity(7));
	}

	public int parity(int num) {
		int count = 0;

		while (num > 0) {
			count += (num & 1);
			num >>= 1;
		}
		if (count % 2 == 0) return 0;
		return 1;
	}
}

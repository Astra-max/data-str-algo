public class SetBits {
	public static void main(String[] args) {
		SetBits st = new SetBits();
		int val = st.countBits(7);
		System.out.printf("set bits in 5: %d\n", val);
	}
	public  int countBits(int num) {
		int count = 0;

		while (num > 0) {
			count += (num & 1);
			num >>= 1;
		}
		return count;
	}
}

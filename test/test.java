import java.util.ArrayList;
import java.util.List;

public class test {
    public static void main(String[] args) {
        int[] nums = {1, 2, 3};
        List<List<Integer>> res = nsolution(nums);
        System.out.println("res = " + res);
    }

    public static List<List<Integer>> nsolution(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        List<Integer> track = new ArrayList<>();
        boolean[] used = new boolean[nums.length];
        backtrack(nums, track, used, res);
        return res;
    }

    private static void backtrack(int[] nums, List<Integer> track, boolean[] used, List<List<Integer>> res) {
        if (track.size() == nums.length) {
            System.out.println("track = " + track);
            res.add(new ArrayList<>(track));
            return;
        }

        for (int i = 0; i < nums.length; i++) {
            if (used[i]) {
                continue;
            }
            used[i] = true;
            track.add(nums[i]);
            backtrack(nums, track, used, res);
            used[i] = false;
            track.remove(track.size() - 1);
        }
    }
}

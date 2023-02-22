#include <iostream>
#include <vector>
#include <map>

using namespace std;

class Solution {

    public:
        int rob(vector<int>& nums) {

            int sum = 0;
            map <int, bool> mapp;
            map<int, bool>::iterator itr;


            for (int i = 0; i < (nums.size()-1); i++) {
                int max = 0, j = 0;
                for (;j < (nums.size()); j++) {
                    cout << nums[j] << endl;
                    if ((nums[j] > max) && !mapp[j]) {
                        cout << nums[j] <<endl;
                        max = nums[j];
                    }
                }

                mapp[j-1] = true; mapp[j] = true; mapp[j+1] = true;
                for (itr = mapp.begin(); itr != mapp.end(); ++itr) {
                        cout << itr->first << ": "<< itr ->second << endl;
                }
                sum += max;
            }
            return sum;
        }
};

int main() {
    vector <int> numbers = {4, 2, 4, 4};
    map <int, int> isgod;
    Solution solution;
    int max =solution.rob(numbers);
    cout <<"max: " << max << endl;

    return 0;
}

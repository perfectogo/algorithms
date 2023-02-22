#include <iostream>

class Solution {
    public:
        int sum;
        
        bool isHappy(int n) {

            if (rec(n) == 1) {

                return 1;
            }

            return 0;
        }

        int rec (int n) {
            while (n >= 1) {
                sum  = (n % 10 ) * (n % 10 );
                n /= 10;
            }

            std::cout << sum << std::endl;

            if (sum < 7) {
                return sum;
            }

            return rec(sum);
        }
} s;

int main() {
    std::cout << "go\n";
    std::cout << s.isHappy(7) << std::endl;
}
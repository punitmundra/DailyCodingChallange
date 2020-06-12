/*
26. Remove Duplicates from Sorted Array
*/
 #include <bits/stdc++.h> 
 using namespace std;
 void removeDuplicates(vector<int>& nums) {
     int i =0;
     for(int j = 1 ; j < nums.size();j++){
         if (nums[j] != nums[i]){
             i++;
         }  
      nums[i]=nums[j];
     }
     for (auto a : nums )
        cout << a <<" ";
 }
int main(){
    vector<int> v = {1,2,2,3,3,3,5,6,6,4,7,8,8,8,9,10,10,11,12,12,13,14,14,15,15,16,16,16,16,17};
    removeDuplicates(v);
}

/*
994. Rotting Oranges
In a given grid, each cell can have one of three values:
the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.
Return the minimum number of minutes that must elapse until no cell has a fresh orange.  
If this is impossible, return -1 instead.
*/

#include<iostream> 
#include<cstring> 
#include<bits/stdc++.h> 
#include<queue>
using namespace std; 

pair<pair<int, int>,int> p;

int orangesRotting(vector<vector<int>> &grid)
{
    int row = grid.size();
    if (row == 0)
        return 0;
    int col = grid[0].size();
    int numOf1=0,numOf2=0;
    queue<pair<int,int>> q;
    //cout << "row:"<<row<<" col:"<<col<<endl;    
    for (int a =0;a<row;a++){
        for (int aa=0; aa < grid[a].size(); aa++) {
            if (grid[a][aa] == 1)
                numOf1 +=1;
            if (grid[a][aa] == 2){
                numOf2 +=1;    
                q.push(make_pair(a,aa));
            }
        }
    }
    if (numOf1 == 0)
        return 0; //    
    if (numOf2 == 0) 
        return -1;
    //cout << "Fresh : "<<numOf1 << " rotten:"<<numOf2<<endl;
    
    int day = 0;
    while(!q.empty() and numOf1 > 0){
        int qSize = q.size();
        //cout << "qSize:"<<qSize<<endl;   
        for(int i = 0;i<qSize;i++) {
            auto orange = q.front();
            q.pop();
            /*
            cout<<orange.first<<"  "<<orange.second<<endl;
            cout << orange.first + 1 << "  " << orange.second << endl;
            cout << orange.first - 1 << "  " << orange.second << endl;
            cout << orange.first << "  " << orange.second + 1 << endl;
            cout << orange.first << "  " << orange.second - 1 << endl;
            */
            if ((orange.first + 1) < row && grid[orange.first +1][orange.second]==1) { //next in same row
                q.push(make_pair(orange.first + 1,orange.second));
                grid[orange.first +1][orange.second]=2;
                numOf1--;
            }
            if (0 <= (orange.first - 1) && grid[orange.first - 1][orange.second] == 1)
            { //prev in same row
                q.push(make_pair(orange.first - 1, orange.second));
                grid[orange.first - 1][orange.second] = 2;
                numOf1--;
            }
            if ((orange.second+1 < col) && grid[orange.first][orange.second+1] == 1){
                q.push(make_pair(orange.first, orange.second+1));
                grid[orange.first][orange.second+1] = 2 ;
                numOf1--;
            }
            if ((0 <= orange.second-1) && grid[orange.first][orange.second-1] == 1){
                q.push(make_pair(orange.first, orange.second-1));
                grid[orange.first][orange.second-1] = 2;
                numOf1--;
            }
        }// end for : that day

        day++;
    }
        //cout <<" numOf1:"<<numOf1<<endl;
        if (numOf1 != 0)
            return -1;
        else
            return day;
}

int main(){
    vector<int> v1 {0, 0, 0, 1, 1};
    vector<int> v2 {0, 0, 1, 1, 2};
    vector<int> v3 {0, 1, 1, 0, 0};
    vector<int> v4 {1, 1, 0, 0, 0};
    vector<int> v5 {2, 0, 0, 0, 0};

    vector<vector<int>> v;
    v.push_back(v1);
    v.push_back(v2);
    v.push_back(v3);
    v.push_back(v4);
    v.push_back(v5);
    cout<<orangesRotting(v);
     return 0;
}

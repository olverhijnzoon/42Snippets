#include <stdio.h>
#include <time.h>
#include <stdlib.h>

// Binary search function with O(log N)
int binary_search(int array[], int left_index, int right_index, int target_value) 
{ 
    // "left" represents the leftmost index of the segment of the array
    while (left_index <= right_index) { 
        // Figure out the middle_index by integer division e.g. 7 / 2 = 3
        int middle_index = left_index + (right_index - left_index) / 2; 
  
        // Check if target_value is present at middle_index 
        if (array[middle_index] == target_value) 
            return middle_index; 
  
        // If target_value is greater, ignore the left half 
        if (array[middle_index] < target_value) 
            left_index = middle_index + 1; 
  
        // If target_value is smaller, ignore the right half 
        else
            right_index = middle_index - 1; 
    } 
  
    // if we reach here, the element was not present 
    return -1; 
} 

int main(){
    printf("42 Snippets - Time Complexity\n");

    // binary search algorithmoperates on sorted arrays and has a time complexity of O(log⁡ N)
    printf("Binary Search\n");

    const double TARGET_TIMES[5] = {5, 6, 7, 8, 9}; // Five target times in seconds
    const int INCREMENT = 1 << 22;   // Increment size by this amount. You can adjust it.
    const int M = 10000000;         // Number of searches

    for (int t = 0; t < 5; t++) {
        int N = INCREMENT;
        while (1) {
            int *array = (int*) malloc(N * sizeof(int));
            for (int i = 0; i < N; i++)
                array[i] = i;

            clock_t start, end;
            double cpu_time_used;

            start = clock();
            
            for (int i = 0; i < M; i++) {
                binary_search(array, 0, N - 1, rand() % N);
            }

            end = clock();
            cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
            
            free(array);  // Free the allocated memory

            printf("Tried target time of %fs with N %d\n", TARGET_TIMES[t], N);
            
            if (cpu_time_used > TARGET_TIMES[t]) {
                printf("For target time of %fs, maximum N is %d\n", TARGET_TIMES[t], N);
                break;
            }

            N += INCREMENT;  // Increase the size and repeat
        }
    }

    return 0;
}

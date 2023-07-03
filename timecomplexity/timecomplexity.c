#include <stdio.h>
#include <time.h>
#include <stdlib.h>

// Linear search function with O(N)
int linear_search(int arr[], int n, int x)
{
   for (int i = 0; i < n; i++)
       if (arr[i] == x)
           return i;
   return -1;
}

// Bubble sort function with O(N^2)
void bubble_sort(int arr[], int n)
{
   int i, j, temp;
   for (i = 0; i < n-1; i++)      
       for (j = 0; j < n-i-1; j++) 
           if (arr[j] > arr[j+1])
           {
              temp = arr[j];
              arr[j] = arr[j+1];
              arr[j+1] = temp;
           }
}

int main()
{
    int N_values[5] = {100000, 200000, 300000, 400000, 500000};
    int M = 100000;  // Number of times to perform linear search

    for (int n = 0; n < 5; n++) {
        int N = N_values[n];
        int *arr = (int*) malloc(N * sizeof(int));  // Dynamic allocation
        
        // Fill array with nearly sorted values
        for (int i = 0; i < N; i++)
            arr[i] = i;

        // Introduce some randomness
        for (int i = 0; i < N / 1000; i++) {
            int j = rand() % N;
            int k = rand() % N;
            int temp = arr[j];
            arr[j] = arr[k];
            arr[k] = temp;
        }

        clock_t start, end;
        double cpu_time_used;
         
        start = clock();

        // Perform linear search with handicap
        for (int i = 0; i < M; i++){
            linear_search(arr, N, rand() % N);
        }

        end = clock();
        cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
        printf("N=%d \t %fs\t ", N, cpu_time_used);
        
        start = clock();

        // Perform bubble sort
        bubble_sort(arr, N);
        
        end = clock();
        cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
        printf("%fs\n", cpu_time_used);

        free(arr);  // Free the allocated memory
    }

    return 0;
}

#include <iostream>

// CUDAカーネル関数
__global__ void helloWorldKernel() {
    printf("Hello, World from GPU!\n");
}

int main() {
    // GPU上でカーネルを実行
    helloWorldKernel<<<10, 10>>>();

    // CUDAのデバイス同期を待機
    cudaDeviceSynchronize();

    // ホスト側のメッセージ
    std::cout << "Hello, World from CPU!" << std::endl;

    return 0;
}
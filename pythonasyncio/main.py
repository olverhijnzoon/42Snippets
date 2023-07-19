import asyncio

async def task(name, time):
    print(f'Task {name} started')
    await asyncio.sleep(time)
    print(f'Task {name} completed after {time} seconds')

async def main():
    task1 = asyncio.create_task(task("A", 2))
    task2 = asyncio.create_task(task("B", 1))
    task3 = asyncio.create_task(task("C", 3))

    await task1
    await task2
    await task3

if __name__ == "__main__":
    asyncio.run(main())

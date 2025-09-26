import statistics
from concurrent.futures import ProcessPoolExecutor as Pool
import random
import csv


def create_scv():
    categories = "ABCD"
    for i in range(5):
        with open(f"csv_{i + 1}.csv", "w", newline="") as csvfile:
            writer = csv.writer(csvfile)
            for j in range(50):
                writer.writerow([random.choice(categories), random.random()])
    print("5 CSV создано\n")


def compute_median_deviation(file):
    #         A   B   C   D
    values = [[], [], [], []]
    with open(file) as f:
        reader = csv.reader(f)
        for r in reader:
            if len(r) != 2:
                continue
            category, value = r[0], float(r[1])
            if category == 'A':
                values[0].append(value)
            elif category == 'B':
                values[1].append(value)
            elif category == 'C':
                values[2].append(value)
            elif category == 'D':
                values[3].append(value)
    res = []
    for v in values:
        median = statistics.median(v)
        deviation = statistics.stdev(v)
        res.append((median, deviation))
    return res


def processing_csv():
    files = [f"csv_{i + 1}.csv" for i in range(5)]
    with Pool() as pool:
        res = list(pool.map(compute_median_deviation, files))
    #res = list(map(compute_median_deviation, files))
    categories_medians = [[], [], [], []]
    categories_deviations = [[], [], [], []]
    for r in res:
        for i, (median, deviation) in enumerate(r):
            categories_medians[i].append(median)
            categories_deviations[i].append(deviation)

    overall_medians = [statistics.median(medians) for medians in categories_medians]
    median_deviations = [statistics.stdev(medians) for medians in categories_medians]

    categories = "ABCD"
    for i, file in enumerate(files):
        print(f"Файл {file}:")
        for j, category in enumerate(categories):
            median, dev = res[i][j]
            print(f"{category}: {median}, {dev}")
        print()

    for j, category in enumerate(categories):
        print(f"Медиана медиан {category}: {overall_medians[j]}, Среднее отклонение медиан {category}: {median_deviations[j]}")
    print()

def main():
    create_scv()
    processing_csv()


if __name__ == "__main__":
    main()

# Лаб_1

# Необходимо на python написать код генерации 5 csv файлов вида:

# Категория, значение. Где категория это рандомная буква от A до D, а значение рандомное число с плавающей точкой (float к примеру)

# Далее нужно параллельно обработать эти 5 файлов. Обработка заключается в поиске медианы и стандартного отклонения в рамках каждой буквы. Т.е. на выходе вы получите результат в виде:

# А, медиана, отклонение
# Б, медиана, отклонение
# ...
# Д, медиана, отклонение

# После этого нужно из результатов также найти медиану из медиан и стандартное отклонение из медиан. Получив такой же результат

# А, медиана медиан, стандартное отклонение медиан
# ...
# Д, то же самое

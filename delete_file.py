import os

def list_files_in_directory(path):
    # 获取当前文件夹中的所有文件和文件夹
    items = os.listdir(path)

    for item in items:
        item_path = os.path.join(path, item)

        # 如果是文件，则打印文件路径
        if os.path.isfile(item_path):
            if ".DS_Store" in item_path:
                os.remove(item_path)
                print("File:", item_path, "delete complete!")
        
        # 如果是文件夹，则递归调用该函数
        elif os.path.isdir(item_path):
            list_files_in_directory(item_path)

        
if __name__ == '__main__':
    # 指定要查看的文件夹路径
    directory_path = "./"

    # 调用函数开始遍历文件夹
    list_files_in_directory(directory_path)
    
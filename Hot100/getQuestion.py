import os
import json
import random

# 获取脚本所在的绝对路径
SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))

QUESTION_FILE = 'hot100.json'
SOLVED_FILE = 'solved'

def getQuestionList(filename):
    try:
        # 构建绝对路径
        abs_path = os.path.join(SCRIPT_DIR, filename)
        
        # 确保目标目录存在
        os.makedirs(os.path.dirname(abs_path), exist_ok=True)

        with open(abs_path, 'r', encoding='utf-8') as file:
            questionList = json.load(file)
        # 输出读取到的题目列表
        return questionList
    except IOError as e:
        print(f"Error reading file {abs_path}: {str(e)}")
        return False
    except Exception as e:
        print(f"Unexpected error while reading file: {str(e)}")
        return False

def getSolvedList(filename):
    try:
         # 构建绝对路径
        abs_path = os.path.join(SCRIPT_DIR, filename)
        
        # 确保目标目录存在
        os.makedirs(os.path.dirname(abs_path), exist_ok=True)

        with open(abs_path, 'a+', encoding='utf-8') as file:
            file.seek(0)
            lines = file.readlines()
            # 将每行转换为整数并存储到列表中
            solved_list = [int(line) for line in lines]
        return solved_list
    except IOError as e:
        print(f"Error reading file {abs_path}: {str(e)}")
        return False
    except Exception as e:
        print(f"Unexpected error while reading file: {str(e)}")
        return False

def chooseQuestion(questions, solved):
    randNum = random.randint(0, 99)
    question = questions[randNum]
    while question['id'] in solved:
        randNum = random.randint(0, 99)
        question = questions[randNum]
    # print(f"quesiont id: {int(question['id'])}")
    # print(solved)
    return question

def printQuestion(question):
    print('已为您选择题目，相关信息如下:')
    print(f"题目ID: {int(question['id'])}\t题目: {str(question['title'])}\t难度: {str(question['difficulty'])}")
    print(f"中文Leetcode: https://leetcode.cn/problems/{str(question['en_name'])}")
    print(f"Leetcode: https://leetcode.com/problems/{str(question['en_name'])}")

def solveQuestion(id, filename):
    try:
         # 构建绝对路径
        abs_path = os.path.join(SCRIPT_DIR, filename)
        
        # 确保目标目录存在
        os.makedirs(os.path.dirname(abs_path), exist_ok=True)

        with open(abs_path, 'a+', encoding='utf-8') as file:
            file.write(str(id)+'\n')
    except IOError as e:
        print(f"Error writing to file {abs_path}: {str(e)}")
        return False
    except Exception as e:
        print(f"Unexpected error while writing to file: {str(e)}")
        return False

def clearSolvedList(filename):
    try:
         # 构建绝对路径
        abs_path = os.path.join(SCRIPT_DIR, filename)
        
        # 确保目标目录存在
        os.makedirs(os.path.dirname(abs_path), exist_ok=True)

        with open(abs_path, 'w', encoding='utf-8') as file:
            pass # 直接打开并关闭文件将清空文件内容
    except IOError as e:
        print(f"Error clearing file {abs_path}: {str(e)}")
        return False
    except Exception as e:
        print(f"Unexpected error while clearing file: {str(e)}")
        return False

def main():
    questions = getQuestionList(QUESTION_FILE)
    solved = getSolvedList(SOLVED_FILE)
    if len(solved) > 99: 
        # 完成全部的 Hot100，重新生成 solved 文件
        print("已完成全部题目, 正在为您清空刷题记录")
        clearSolvedList(SOLVED_FILE)
        solved = getSolvedList(SOLVED_FILE)

    if questions:
        # 选择一道没被解决的题目
        question = chooseQuestion(questions, solved)
        # 打印题目相关信息
        printQuestion(question)
        # 题目标记为已解决
        solveQuestion(question['id'], SOLVED_FILE)


# 添加这一行来执行 main 函数
if __name__ == "__main__":
    main()
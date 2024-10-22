from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import WebDriverException
import os
import re
import json

# 获取脚本所在的绝对路径
SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))

def setup_chrome_driver():
    """Set up Chrome driver with proper options and error handling."""
    options = webdriver.ChromeOptions()
    options.add_argument('--headless')
    options.add_argument('--disable-gpu')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')
    
    # Add user agent to avoid detection
    options.add_argument('--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36')
    
    # 添加 Arch 的 chromedriver
    service = Service('/usr/bin/chromedriver')
    try:
        # Try using the system Chrome driver
        driver = webdriver.Chrome(options=options, service=service)
        return driver
    except WebDriverException as e:
        print("Error setting up Chrome driver:", str(e))
        print("\nTrying to install required components...")

def get_dynamic_content(url, max_retries=3):
    """Get dynamic content from URL with retry mechanism."""
    
    for attempt in range(max_retries):
        driver = setup_chrome_driver()
        if not driver:
            return None
            
        try:
            print(f"Attempt {attempt + 1}/{max_retries} to fetch content...")
            driver.get(url)
            
            # Wait for content to load
            element = WebDriverWait(driver, 15).until(
                EC.presence_of_element_located((By.XPATH, 
                    "/html/body/div[1]/div[1]/div[5]/div/div/div[2]/div[2]/div/div/div[101]/div"))
            )
            
            # 找到包含 data-rbd-draggable-id 的所有 div 元素
            div_elements = element.find_elements(By.XPATH, "//div[@data-rbd-draggable-id]")
            # 提取每个 div 的 data-rbd-draggable-id 属性
            draggable_ids = [div.get_attribute('data-rbd-draggable-id') for div in div_elements]

            # Get rendered content
            content = element.text

            return content, draggable_ids
            
        except Exception as e:
            print(f"Error during attempt {attempt + 1}: {str(e)}")
            if attempt < max_retries - 1:
                print("Retrying...")
            
        finally:
            driver.quit()
    
    return None

def organize_hot100_message(content, ids):
    # print(content)
    # 使用正则表达式匹配出题号、题目名称和难度
    pattern = re.compile(r'(\d+)\.\s(.*?)\n(简单|中等|困难)')

    # 使用 findall 找出所有符合的匹配项
    matches = pattern.findall(content)

    # 将匹配结果整理为一个列表
    problems = [{"id": int(match[0]), "title": match[1], 
                 "difficulty": match[2], "en_name": ids[index]} 
                 for index, match in enumerate(matches)]

    # 返回整理后的结构
    return problems

# 保存爬取到的 Hot100 题目信息
def save_hot100_content(content, filename):
    # 直接保存即可
    try:
        # 构建绝对路径
        abs_path = os.path.join(SCRIPT_DIR, filename)
        
        # 确保目标目录存在
        os.makedirs(os.path.dirname(abs_path), exist_ok=True)
        
        # 写入文件
        with open(abs_path, "w", encoding="utf-8") as file:
            json.dump(content, file, ensure_ascii=False, indent=4)
        print(f"Successfully saved content to {abs_path}")
        return True
        
    except IOError as e:
        print(f"Error writing to file {abs_path}: {str(e)}")
        return False
    except Exception as e:
        print(f"Unexpected error while writing to file: {str(e)}")
        return False

def main():
    # baseurl = 'https://leetcode.cn/'
    extraurl = 'https://leetcode.cn/problem-list/2cktkvj/'
    filename = 'hot100.json'
    
    content, ids = get_dynamic_content(extraurl)
    # print(ids)
    if content:
        print("Successfully retrieved content")
        print("Generating Hot100 List")
        problems = organize_hot100_message(content, ids)
        print("Saving Hot100 List")
        save_hot100_content(content=problems, filename=filename)
    else:
        print("Failed to retrieve content after all attempts.")

if __name__ == "__main__":
    main()
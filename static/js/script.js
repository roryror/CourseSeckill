document.addEventListener('DOMContentLoaded', function() {
    // 获取DOM元素
    const coursesTable = document.getElementById('courses-table').getElementsByTagName('tbody')[0];
    const cacheTable = document.getElementById('cache-table').getElementsByTagName('tbody')[0];
    const ordersTable = document.getElementById('orders-table').getElementsByTagName('tbody')[0];
    const showStockBtn = document.getElementById('show-stock-btn');
    const showOrderStatusBtn = document.getElementById('show-order-status-btn');
    const showOrdersBtn = document.getElementById('show-orders-btn');
    const seckillBtn = document.getElementById('seckill-btn');
    const userIdInput = document.getElementById('user-id');
    const warmupBtn = document.getElementById('warmup-btn');
    const warmupMessage = document.getElementById('warmup-message');

    // 页面加载时自动获取课程列表
    fetchCourses();

    // 绑定按钮点击事件
    showStockBtn.addEventListener('click', fetchStockCache);
    showOrderStatusBtn.addEventListener('click', fetchOrderStatusCache);
    showOrdersBtn.addEventListener('click', fetchOrders);
    seckillBtn.addEventListener('click', seckillSelectedCourses);
    warmupBtn.addEventListener('click', warmupSystem);

    // 系统预热
    function warmupSystem() {
        warmupBtn.disabled = true;
        warmupBtn.textContent = '预热中...';
        warmupMessage.style.display = 'block';
        warmupMessage.className = 'message';
        warmupMessage.textContent = '正在预热系统...';
        
        fetch('/warmup/', {
            method: 'POST'
        })
        .then(response => response.json())
        .then(data => {
            warmupMessage.className = 'success-message';
            warmupMessage.textContent = '系统预热成功！';
            
            // 1秒后自动展示库存缓存
            setTimeout(() => {
                fetchStockCache();
                
                // 再过2秒后隐藏成功消息
                setTimeout(() => {
                    warmupMessage.style.display = 'none';
                }, 2000);
            }, 1000);
        })
        .catch(error => {
            console.error('Error:', error);
            warmupMessage.className = 'error-message';
            warmupMessage.textContent = '系统预热失败，请稍后重试！';
            
            // 3秒后隐藏错误消息
            setTimeout(() => {
                warmupMessage.style.display = 'none';
            }, 3000);
        })
        .finally(() => {
            warmupBtn.disabled = false;
            warmupBtn.textContent = '系统预热';
        });
    }

    // 获取课程列表
    function fetchCourses() {
        showLoading(coursesTable);
        fetch('/viewer/courses')
            .then(response => {
                if (!response.ok) {
                    throw new Error('获取课程列表失败');
                }
                return response.json();
            })
            .then(data => {
                clearTable(coursesTable);
                if (data.courses && data.courses.length > 0) {
                    data.courses.forEach(course => {
                        const row = coursesTable.insertRow();
                        
                        // 添加复选框
                        const checkboxCell = row.insertCell(0);
                        checkboxCell.className = 'checkbox-cell';
                        const checkbox = document.createElement('input');
                        checkbox.type = 'checkbox';
                        checkbox.dataset.courseId = course.ID;
                        checkboxCell.appendChild(checkbox);
                        
                        // 添加课程信息
                        row.insertCell(1).textContent = course.ID;
                        row.insertCell(2).textContent = course.Name;
                        row.insertCell(3).textContent = course.Stock;
                        row.insertCell(4).textContent = course.MaxStock;
                        row.insertCell(5).textContent = course.MinStock;
                    });
                } else {
                    showNoData(coursesTable, 5);
                }
            })
            .catch(error => {
                console.error('Error:', error);
                showError(coursesTable, 5, '获取课程列表失败');
            });
    }

    // 获取库存缓存
    function fetchStockCache() {
        showLoading(cacheTable);
        fetch('/viewer/stock')
            .then(response => {
                if (!response.ok) {
                    throw new Error('获取库存缓存失败');
                }
                return response.json();
            })
            .then(data => {
                clearTable(cacheTable);
                if (data.stock && Object.keys(data.stock).length > 0) {
                    for (const [key, value] of Object.entries(data.stock)) {
                        const row = cacheTable.insertRow();
                        row.insertCell(0).textContent = `课程ID: ${key}`;
                        row.insertCell(1).textContent = `库存: ${value}`;
                    }
                } else {
                    showNoData(cacheTable, 2);
                }
            })
            .catch(error => {
                console.error('Error:', error);
                showError(cacheTable, 2, '获取库存缓存失败');
            });
    }

    // 获取订单状态缓存
    function fetchOrderStatusCache() {
        showLoading(cacheTable);
        fetch('/viewer/order-status')
            .then(response => {
                if (!response.ok) {
                    throw new Error('获取订单状态缓存失败');
                }
                return response.json();
            })
            .then(data => {
                clearTable(cacheTable);
                if (data.orderStatus && Object.keys(data.orderStatus).length > 0) {
                    for (const [key, value] of Object.entries(data.orderStatus)) {
                        const row = cacheTable.insertRow();
                        row.insertCell(0).textContent = key;
                        row.insertCell(1).textContent = value;
                    }
                } else {
                    showNoData(cacheTable, 2);
                }
            })
            .catch(error => {
                console.error('Error:', error);
                showError(cacheTable, 2, '获取订单状态缓存失败');
            });
    }

    // 获取订单列表
    function fetchOrders() {
        showLoading(ordersTable);
        fetch('/viewer/orders')
            .then(response => {
                if (!response.ok) {
                    throw new Error('获取订单列表失败');
                }
                return response.json();
            })
            .then(data => {
                clearTable(ordersTable);
                if (data.orders && data.orders.length > 0) {
                    data.orders.forEach(order => {
                        const row = ordersTable.insertRow();
                        row.insertCell(0).textContent = order.ID;
                        row.insertCell(1).textContent = order.UserID;
                        row.insertCell(2).textContent = order.CourseID;
                    });
                } else {
                    showNoData(ordersTable, 3);
                }
            })
            .catch(error => {
                console.error('Error:', error);
                showError(ordersTable, 3, '获取订单列表失败');
            });
    }

    // 秒杀选中的课程
    function seckillSelectedCourses() {
        const userId = userIdInput.value;
        if (!userId) {
            alert('请输入用户ID');
            return;
        }

        const checkboxes = coursesTable.querySelectorAll('input[type="checkbox"]:checked');
        if (checkboxes.length === 0) {
            alert('请选择至少一个课程');
            return;
        }

        // 按照从上到下的顺序获取选中的课程ID
        const selectedCourseIds = [];
        coursesTable.querySelectorAll('input[type="checkbox"]').forEach(checkbox => {
            if (checkbox.checked) {
                selectedCourseIds.push(checkbox.dataset.courseId);
            }
        });

        // 依次秒杀选中的课程
        seckillCourseSequentially(userId, selectedCourseIds, 0);
    }

    // 按顺序秒杀课程
    function seckillCourseSequentially(userId, courseIds, index) {
        if (index >= courseIds.length) {
            // 所有课程都已秒杀完成，刷新数据
            fetchCourses();
            return;
        }

        const courseId = courseIds[index];
        
        // 发送秒杀请求
        fetch(`/seckill/${courseId}/${userId}`, {
            method: 'POST'
        })
        .then(response => response.json())
        .then(data => {
            // 显示秒杀结果
            const message = document.createElement('div');
            message.className = data.error ? 'error-message' : 'success-message';
            message.textContent = `课程ID ${courseId}: ${data.message}`;
            
            // 将消息插入到秒杀按钮后面
            const actionsDiv = document.querySelector('.section:first-child .actions');
            actionsDiv.appendChild(message);
            
            // 3秒后移除消息
            setTimeout(() => {
                message.remove();
            }, 3000);
            
            // 继续秒杀下一个课程
            setTimeout(() => {
                seckillCourseSequentially(userId, courseIds, index + 1);
            }, 500);
        })
        .catch(error => {
            console.error('Error:', error);
            
            // 显示错误消息
            const message = document.createElement('div');
            message.className = 'error-message';
            message.textContent = `课程ID ${courseId}: 秒杀请求失败`;
            
            const actionsDiv = document.querySelector('.section:first-child .actions');
            actionsDiv.appendChild(message);
            
            // 3秒后移除消息
            setTimeout(() => {
                message.remove();
            }, 3000);
            
            // 继续秒杀下一个课程
            setTimeout(() => {
                seckillCourseSequentially(userId, courseIds, index + 1);
            }, 500);
        });
    }

    // 辅助函数：清空表格
    function clearTable(table) {
        while (table.firstChild) {
            table.removeChild(table.firstChild);
        }
    }

    // 辅助函数：显示加载中
    function showLoading(table) {
        clearTable(table);
        const row = table.insertRow();
        const cell = row.insertCell(0);
        cell.colSpan = 10; // 足够大的列数
        cell.className = 'loading';
        cell.textContent = '加载中...';
    }

    // 辅助函数：显示无数据
    function showNoData(table, colSpan) {
        clearTable(table);
        const row = table.insertRow();
        const cell = row.insertCell(0);
        cell.colSpan = colSpan;
        cell.className = 'loading';
        cell.textContent = '暂无数据';
    }

    // 辅助函数：显示错误
    function showError(table, colSpan, message) {
        clearTable(table);
        const row = table.insertRow();
        const cell = row.insertCell(0);
        cell.colSpan = colSpan;
        cell.className = 'error-message';
        cell.textContent = message || '加载失败';
    }
}); 
import './App.css';
import 'antd/dist/antd.css';

import React from 'react';
import axios from "axios";
import {Button, Form, Input, message} from 'antd';

export default function App() {

    const onFinish = async (values: any) => {
        try {
            const result = await axios({
                method: "POST",
                url: "/api/login",
                headers: {
                    "Content-Type": "application/json; charset=utf-8"
                },
                data: JSON.stringify(values)
            });
            if (result.status === 200) {
                message.success("Good Job, well done! 🥳");
            }
        }
        catch (e) {
            message.error("Ops, you are wrong! 😜");
        }
    };

    const [form] = Form.useForm();

    const onReset = () => {
        form.resetFields();
    };

    return (
        <div id="App">
            <Form
                name="basic"
                form={form}
                labelCol={{ span: 8 }}
                wrapperCol={{ span: 16 }}
                initialValues={{ remember: true }}
                onFinish={onFinish}
                autoComplete="off"
            >
                <Form.Item
                    label="Username"
                    name="username"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Password"
                    name="password"
                    rules={[{ required: true, message: 'Please input your password!' }]}
                >
                    <Input.Password />
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button htmlType="button" onClick={onReset}>
                        Reset
                    </Button>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                </Form.Item>
            </Form>
        </div>
    )
};

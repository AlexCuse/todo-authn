// apiService.js
import {session} from 'keratin-authn';

const API_BASE_URL = 'https://api.todoauthn.com/api/v1';

export async function getTodos() {
    const response = await fetch(`${API_BASE_URL}/todos`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${session()}`,
        },
    });

    if (response.ok) {
        return await response.json();
    } else {
        throw new Error('Failed to fetch todos');
    }
}

export async function addTodo(title) {
    const response = await fetch(`${API_BASE_URL}/todos`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${session()}`,
        },
        body: JSON.stringify({ title }),
    });

    if (response.ok) {
        return await response.json();
    } else {
        throw new Error('Failed to add a todo');
    }
}

export async function updateTodo(id, title, description, completed) {
    const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${session()}`,
        },
        body: JSON.stringify({ title, description, completed }),
    });

    if (response.ok) {
        return await response.json();
    } else {
        throw new Error('Failed to update the todo');
    }
}

export async function deleteTodo(id) {
    const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${session()}`,
        },
    });

    if (response.ok) {
        return true;
    } else {
        throw new Error('Failed to delete the todo');
    }
}

export async function addUser(email) {
    const response = await fetch(`${API_BASE_URL}/user`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${session()}`,
        },
        body: JSON.stringify({ email }),
    });

    if (response.ok) {
        return await response.json();
    } else {
        throw new Error('Failed to add a todo');
    }
}

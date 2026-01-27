import axios from 'axios';
import { UserDataType } from '../../utils/data-types';

export const login = async(data: Pick<UserDataType, 'username' | 'password'>) => {
    const response = await axios.post('http://localhost:3000/auth/login', data, {
        withCredentials: true
    });
    return response;
}

export const register = async(data: UserDataType) => {
    const response = await axios.post('http://localhost:3000/auth/register', data);
    return response;
}

export const logout = async() => {
    const response = await axios.get('http://localhost:3000/auth/logout', {
        withCredentials: true
    });
    return response;
}
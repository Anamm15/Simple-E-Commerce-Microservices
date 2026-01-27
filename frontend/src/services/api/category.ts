import axios from 'axios';

export const getCategories = async() => {
    const response = await axios.get('http://localhost:3000/category/get', {
        withCredentials: true
    })
    return response;
}
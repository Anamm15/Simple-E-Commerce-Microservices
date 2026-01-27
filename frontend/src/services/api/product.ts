import axios from 'axios';
import { ProductDataType } from '../../utils/data-types';

export const getProducts = async() => {
    const response = await axios.get('http://localhost:3000/product/getAllProducts')
    return response;
}

export const getProduct = async(id: number) => {
    return await axios.get('http://localhost:3000/product/get/' + id, {withCredentials: true});
}

export const addProduct = async(data: ProductDataType) => {
    const response = await axios.post('http://localhost:3000/product/add', data, {withCredentials: true});
    return response;
}
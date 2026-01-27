import axios from 'axios';
import { StoreDataType } from '../../utils/data-types';

export const getStoreByUserId = async(userId : number) => {
    return await axios.get('http://localhost:3000/store/getStoreByUserId/' + userId, {
        withCredentials: true});
}

export const addStore = async(data: StoreDataType) => {
    return await axios.post('http://localhost:3000/store/add', data, {withCredentials: true});
}
import axios from 'axios';
import { CommentDataType } from '../../utils/data-types';

export const getCommentsPerProduct = async(id: number) => {
    return await axios.get('http://localhost:3000/comment/getCommentPerProduct/' + id, {withCredentials: true});
}

export const addComment = async(data: CommentDataType) => {
    return await axios.post('http://localhost:3000/comment/add', data, {withCredentials: true});
}
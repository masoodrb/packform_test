import axios from "axios"

const BASE_URL = import.meta.env.VITE_BASE_URL


export const getOrderDetails = (query: string) => {
  let url = `/api/orderDetails?query=${query}`
  if (!query) {
    url = `/api/orderDetails`
  }  
  return axios.get(url)
}
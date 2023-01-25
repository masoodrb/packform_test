import axios from "axios";
import type moment from "moment";

export const getOrderDetails = (query: string) => {
  const url = `/api/order-details`;

  return axios.get(url, {
    params: {
      query,
    },
  });
};

export const filterOrderDetailsByDate = (
  query: string,
  startDate: moment.Moment,
  endDate: moment.Moment
) => {
  const url = `/api/filter-by-date`;
  return axios.get(url, {
    params: {
      query,
      startDate,
      endDate,
    },
  });
};

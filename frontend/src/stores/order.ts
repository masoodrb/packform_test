import { defineStore } from "pinia";
import { getOrderDetails, filterOrderDetailsByDate } from '@/api';
import moment from "moment";

// pinia doesn't have mutations it just have actions that can be used as mutations 
// when they modify the state  

export const useOrderStore = defineStore("order", {

  state: () => ({
    details: [],
    startDate: moment().startOf('month'),
    endDate: moment().endOf('month'),
    query: ""
  }),
  getters: {
    orderDetails: (state) => state.details
  },
  actions: {
    onSearch(value: string) {
      this.query = value
    },
    async loadOrderDetails() {
      const resp = await getOrderDetails(this.query)
      this.details = resp.data
    },
    async dateChanged(value: Array<Date>) {
      this.startDate = moment(value[0])
      this.endDate = moment(value[1])

      const resp = await filterOrderDetailsByDate(this.query, this.startDate, this.endDate)
      this.details = resp.data
    }
  }
});

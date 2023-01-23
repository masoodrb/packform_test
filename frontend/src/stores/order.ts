import { ref, computed } from "vue";
import { defineStore } from "pinia";
import { getOrderDetails } from '@/api';

// pinia doesn't have mutations it just have actions that can be used as mutations 
// when they modify the state  

export const useOrderStore = defineStore("order", {

  state: () => ({
    details: []
  }),
  getters: {
    orderDetails: (state) => state.details
  },
  actions: {
    async loadOrderDetails(query: string = "") {
      const resp = await getOrderDetails(query)
      this.details = resp.data
    }
  }
});

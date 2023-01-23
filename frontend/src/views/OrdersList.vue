<template>
  <div class="container">
    <Header
      @onSearch="onQueryChange"
    >
    </Header>
    <div class="content">
      <div>
        <h3><label>Created date</label></h3>
      </div>
      <div>
        <el-date-picker 
          v-model="selectedDate" 
          type="daterange"
          range-separator="-"
          start-placeholder="Start date"
          end-placeholder="End date"
          @change="onDateChange"
        />
      </div>
      <div>
        <h3>Total Amount: ${{ totalAmount.toFixed(2) }} </h3>
      </div>
      <DataGrid></DataGrid>
    </div>
  </div>
</template>
<script lang="ts" setup>
import moment from 'moment'

import { ref, onMounted, computed } from 'vue';
import { useOrderStore } from '@/stores/order'
import Header from '@/components/Header/index.vue'
import DataGrid from '@/components/DataGrid/index.vue'
import type { OrderDetails } from '@/api';
const store = useOrderStore()

const { loadOrderDetails, dateChanged, onSearch } = useOrderStore()
onMounted(() => {
  loadOrderDetails()
})
const onQueryChange = (value: string) => {
  onSearch(value)
  loadOrderDetails()
}

const totalAmount = computed(() => {
  const sumFunction = (total: number, item: OrderDetails) => {
    let sum = 0.0
    sum = total + (item.PricePerUnit * item.Quantity)
    return sum
  }
  return store.details.reduce(sumFunction, 0)
})

const selectedDate = ref([
  store.startDate,
  store.endDate
]);

const onDateChange = (value: Date[]) => {
  dateChanged(value)
}

</script>

<style lang="scss" scoped>

.container {
  padding: 15px;
}
.content {
  margin: 8px
}
</style>

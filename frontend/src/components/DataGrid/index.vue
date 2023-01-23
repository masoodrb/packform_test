<template lang="">
 <el-row>
      <el-table :data="tableData" :default-sort="{ prop: 'OrderDate', order: 'descending' }" style="width: 100%">
        <el-table-column label="Order name">
          <template v-slot="scope">
            {{ scope.row.OrderName }} <br/> {{ scope.row.ProductName }}
          </template>
        </el-table-column>
        <el-table-column prop="CompanyName" label="Customer Company" />
        <el-table-column prop="CustomerName" label="Customer name" />
        <el-table-column prop="OrderDate" label="Order date" sortable :formatter="formatter">
          <template v-slot="scope">
            {{ moment(scope.row.OrderDate).format('MMM Do, YYYY hh:mm A') }}
          </template>
        </el-table-column>
        <el-table-column label="Delivered amount">
          <template v-slot="scope">
            {{ formatAmount(scope.row.PricePerUnit) }}
          </template>
        </el-table-column>
        <el-table-column label="Total Amount">
          <template v-slot="scope">
            {{ formatAmount(scope.row.PricePerUnit * scope.row.Quantity) }}
          </template>
        </el-table-column>
      </el-table>
    </el-row>
    <el-row>
      <div class="pagination-wrapper">
        <el-pagination :page-sizes="[5, 10]"
          layout="total, sizes, prev, pager, next, jumper" 
          :total="totalRows"
          :page-size="pageSize"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange" />
      </div>
    </el-row>
</template>
<script lang="ts" setup>
import { useOrderStore } from '@/stores/order'
import { computed, ref } from 'vue'
import type { TableColumnCtx } from 'element-plus'
import type { OrderDetails } from '@/api'
import moment from 'moment'
const store = useOrderStore()

const page = ref(1)
const pageSize = ref(5)

const handleSizeChange = (val: number) => {
  pageSize.value = val
}

const formatAmount = (val: number) => {
  if (val == 0) {
    return `-`
  } else {
    return `$${val.toFixed(2)}`
  }
}

const tableData = ref(computed(() => store.details.slice(pageSize.value * page.value - pageSize.value, pageSize.value * page.value)))
const totalRows = computed (() => store.details.length)
const formatter = (row: OrderDetails, column: TableColumnCtx<OrderDetails>) => {
  return row.OrderDate
}

const handleCurrentChange = (val: number) => {
  page.value = val
}

</script>
<style lang="scss">
.pagination-wrapper {
  width: 100%;
  margin-top: 20px;

  .el-pagination {
    justify-content: center;
  }
}
</style>
<template>

  <div class="container">
    <el-row>
      <el-col :md="3" :sm="6" :lg="3">
        <div class="customHeader">
          <div class="searchLabel">
            <font-awesome-icon icon="fa-solid fa-magnifying-glass" class="fa-search-icon" size="3x" />
            <label>Search</label>
          </div>
        </div>
      </el-col>
      <el-col :sm="18" :md="20">
        <el-input>

        </el-input>
      </el-col>
    </el-row>

    <!-- <el-row> -->
    <div>
      <label>Created date</label>
    </div>

    <div>
      <el-date-picker v-model="selectedDate" type="daterange" range-separator="-" start-placeholder="Start date"
        end-placeholder="End date" />
    </div>


    <div>
      Total Amount:
    </div>
    <!-- </el-row> -->


    <el-table :data="tableData" :default-sort="{ prop: 'date', order: 'descending' }" style="width: 100%">

      <el-table-column prop="orderName" label="Order name" />
      <el-table-column prop="companyName" label="Customer Company" />
      <el-table-column prop="customerName" label="Customer name" />
      <el-table-column prop="orderDate" label="Order date" sortable :formatter="formatter" />
      <el-table-column prop="deliveredAmount" label="Delivered amount" />
      <el-table-column prop="totalAmount" label="Total Amount" />

    </el-table>

    <el-row>

      <div class="pagination-wrapper">
        <el-pagination 
          :current-page="1" 
          :page-size="5" 
          :page-sizes="[5, 10]"
          layout="total, sizes, prev, pager, next, jumper" 
          :total="16" 
          @size-change="() => {}"
          @current-change="() => {}"
        />
      </div>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import moment from 'moment'
import type { TableColumnCtx } from 'element-plus'


interface OrderDetails {
  orderName: string,
  companyName: string,
  customerName: string,
  orderDate: Date,
  deliveredAmount: Number,
  totalAmount: Number
}



const formatter = (row: OrderDetails, column: TableColumnCtx<OrderDetails>) => {
  return row.orderDate
}

const selectedDate = [moment().startOf('month').format('YYYY-MM-DD'), moment().endOf('month').format('YYYY-MM-DD')];
const tableData = Array<OrderDetails>();

</script>

<style lang="scss" scoped>
.searchLabel {
  font-weight: 600;

  label {
    font-size: 2em;
    font-weight: bold;
  }

  .fa-search-icon {
    width: 30px;
    height: 30px;
    margin-left: 8px;
    margin-right: 8px;
  }
}

.container {
  padding: 15px;
}

.pagination-wrapper {
  width: 100%;
  margin-top: 20px;

  .el-pagination {
    justify-content: center;
  }
}
</style>

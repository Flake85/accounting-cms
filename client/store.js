import { configureStore } from '@reduxjs/toolkit'
import alertModalReducer from './slices/alertModalSlice'

export default configureStore({
    reducer: {
        alertModal: alertModalReducer,
    },
})
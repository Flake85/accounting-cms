import { createSlice } from '@reduxjs/toolkit'

export const alertModalSlice = createSlice({
    name: 'alertModal',
    initialState: {
        show: false,
        title: '',
        body: '',
    },
    reducers: {
        setAlertData: (state, action) => {
            state.title = action.payload.title
            state.body = action.payload.body
        },
        openAlertModal: (state) => { state.show = true },
        closeAlertModal: (state) => { 
            state.show = false
        }
    },
})

export const { setAlertData, openAlertModal, closeAlertModal } = alertModalSlice.actions

export default alertModalSlice.reducer

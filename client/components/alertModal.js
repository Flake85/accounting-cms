import { useSelector, useDispatch } from 'react-redux'
import { closeAlertModal } from '../slices/alertModalSlice'
import Modal from 'react-bootstrap/Modal'
import Button from 'react-bootstrap/Button'

export function AlertModal() {
    const title = useSelector((state) => state.alertModal.title)
    const body = useSelector((state) => state.alertModal.body)
    const show = useSelector((state) => state.alertModal.show)
    const dispatch = useDispatch()
    return (
        <Modal show={show}
               onHide={() => dispatch(closeAlertModal())}
               backdrop="static"
               keyboard={false} >
            <Modal.Header closeButton>
                <Modal.Title>{ title }</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <p>{ body }</p>
            </Modal.Body>
            <Modal.Footer>
                <Button onClick={() => dispatch(closeAlertModal())}>Close</Button>
            </Modal.Footer>
        </Modal>
    )
}
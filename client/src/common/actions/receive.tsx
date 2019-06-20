export default function receive(type: string, payload: object) {
    return {
        type: type,
        payload: payload,
        receivedAt: Date.now()
    }
}
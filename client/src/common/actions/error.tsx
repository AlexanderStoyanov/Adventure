export default function error(type: string, error: string) {
    return {
        type: type,
        payload: {
            error
        }
    }
}
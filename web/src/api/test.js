import service from '@/utils/request'

export const testT = () => {
    return service({
        url: "/test/testT",
        method: 'post',
    })
}
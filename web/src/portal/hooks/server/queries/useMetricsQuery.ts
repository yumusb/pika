import {useQuery} from '@tanstack/react-query';
import {getAgentMetrics, type GetAgentMetricsRequest} from '@/api/agent';

interface UseMetricsQueryOptions {
    agentId: string;
    type: GetAgentMetricsRequest['type'];
    range?: string;
    start?: number;
    end?: number;
    interfaceName?: string;
}

/**
 * 查询 Agent 历史指标数据
 * @param options 查询选项
 * @returns 历史指标查询结果
 */
export const useMetricsQuery = ({agentId, type, range, start, end, interfaceName}: UseMetricsQueryOptions) => {
    return useQuery({
        queryKey: ['agent', agentId, 'metrics', type, range, start, end, interfaceName],
        queryFn: () =>
            getAgentMetrics({
                agentId,
                type,
                range,
                start,
                end,
                interface: interfaceName,
            }),
        enabled: !!agentId,
        // refetchInterval: 30000, // 30秒自动刷新
    });
};

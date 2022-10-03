package scheduler

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/G-Research/armada/internal/scheduler/schedulerobjects"
)

func TestSelectNodeForPod(t *testing.T) {
	type ReqWithExpectation struct {
		Req *schedulerobjects.PodRequirements
		// Whether we expect to find a node for this pod.
		ExpectSuccess bool
	}
	tests := map[string]struct {
		Nodes []*SchedulerNode
		Reqs  []*ReqWithExpectation
	}{
		"cpu 1": {
			Nodes: testNodeItems1,
			Reqs: []*ReqWithExpectation{
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("1"),
							},
						},
					},
					ExpectSuccess: true,
				},
			},
		},
		"cpu 7": {
			Nodes: testNodeItems1,
			Reqs: []*ReqWithExpectation{
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("7"),
							},
						},
					},
					ExpectSuccess: true,
				},
			},
		},
		"cpu 8": {
			Nodes: testNodeItems1,
			Reqs: []*ReqWithExpectation{
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("8"),
							},
						},
					},
					ExpectSuccess: false,
				},
			},
		},
		"all cpu at priority 0": {
			Nodes: testNodeItems1,
			Reqs: []*ReqWithExpectation{
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("7"),
							},
						},
					},
					ExpectSuccess: true,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("4"),
							},
						},
					},
					ExpectSuccess: true,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("1"),
							},
						},
					},
					ExpectSuccess: true,
				},
			},
		},
		"running total": {
			Nodes: testNodeItems1,
			Reqs: []*ReqWithExpectation{
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("7"),
							},
						},
					},
					ExpectSuccess: true,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("5"),
							},
						},
					},
					ExpectSuccess: false,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("4"),
							},
						},
					},
					ExpectSuccess: true,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("2"),
							},
						},
					},
					ExpectSuccess: false,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("1"),
							},
						},
					},
					ExpectSuccess: true,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu": resource.MustParse("1"),
							},
						},
					},
					ExpectSuccess: false,
				},
			},
		},
		"running total with memory": {
			Nodes: testNodeItems1,
			Reqs: []*ReqWithExpectation{
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu":    resource.MustParse("7"),
								"memory": resource.MustParse("7Gi"),
							},
						},
					},
					ExpectSuccess: true,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu":    resource.MustParse("5"),
								"memory": resource.MustParse("5Gi"),
							},
						},
					},
					ExpectSuccess: false,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu":    resource.MustParse("4"),
								"memory": resource.MustParse("4Gi"),
							},
						},
					},
					ExpectSuccess: true,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu":    resource.MustParse("2"),
								"memory": resource.MustParse("2Gi"),
							},
						},
					},
					ExpectSuccess: false,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu":    resource.MustParse("1"),
								"memory": resource.MustParse("1Gi"),
							},
						},
					},
					ExpectSuccess: true,
				},
				{
					Req: &schedulerobjects.PodRequirements{
						Priority: 0,
						ResourceRequirements: &v1.ResourceRequirements{
							Requests: v1.ResourceList{
								"cpu":    resource.MustParse("1"),
								"memory": resource.MustParse("1Gi"),
							},
						},
					},
					ExpectSuccess: false,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			db, err := NewNodeDb(testPriorities, testResources)
			if !assert.NoError(t, err) {
				return
			}
			err = db.Upsert(tc.Nodes)
			if !assert.NoError(t, err) {
				return
			}

			for _, req := range tc.Reqs {
				report, err := db.SelectAndBindNodeToPod(uuid.New(), req.Req)
				if !assert.NoError(t, err) {
					return
				}
				if !assert.NotNil(t, report) {
					return
				}
				fmt.Println(report.String())
				if req.ExpectSuccess {
					assert.NotNil(t, report.Node)
				} else {
					assert.Nil(t, report.Node)
				}
			}
		})
	}
}

func benchmarkUpsert(numNodes int, b *testing.B) {
	db, err := NewNodeDb(testPriorities, testResources)
	if !assert.NoError(b, err) {
		return
	}
	nodes := testNodeItems2(testPriorities, testResources, numNodes)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := db.Upsert(nodes)
		if !assert.NoError(b, err) {
			return
		}
	}
}

func BenchmarkUpsert1(b *testing.B)      { benchmarkUpsert(1, b) }
func BenchmarkUpsert1000(b *testing.B)   { benchmarkUpsert(1000, b) }
func BenchmarkUpsert100000(b *testing.B) { benchmarkUpsert(100000, b) }

func benchmarkSelectAndBindNodeToPod(
	numCpuNodes, numTaintedCpuNodes, numGpuNodes,
	numSmallCpuJobsToSchedule, numLargeCpuJobsToSchedule, numGpuJobsToSchedule int,
	b *testing.B,
) {
	db, err := NewNodeDb(testPriorities, testResources)
	if !assert.NoError(b, err) {
		return
	}
	nodes := testNodes3(numCpuNodes, numTaintedCpuNodes, numGpuNodes, testPriorities)
	err = db.Upsert(nodes)
	if !assert.NoError(b, err) {
		return
	}

	smallCpuJob := &schedulerobjects.PodRequirements{
		Priority: 0,
		ResourceRequirements: &v1.ResourceRequirements{
			Requests: v1.ResourceList{
				"cpu":    resource.MustParse("1"),
				"memory": resource.MustParse("4Gi"),
			},
		},
	}
	largeCpuJob := &schedulerobjects.PodRequirements{
		Priority: 0,
		ResourceRequirements: &v1.ResourceRequirements{
			Requests: v1.ResourceList{
				"cpu":    resource.MustParse("32"),
				"memory": resource.MustParse("256Gi"),
			},
		},
		Tolerations: []v1.Toleration{
			{
				Key:   "largeJobsOnly",
				Value: "true",
			},
		},
	}
	gpuJob := &schedulerobjects.PodRequirements{
		Priority: 0,
		ResourceRequirements: &v1.ResourceRequirements{
			Requests: v1.ResourceList{
				"cpu":    resource.MustParse("4"),
				"memory": resource.MustParse("16Gi"),
				"gpu":    resource.MustParse("1"),
			},
		},
		Tolerations: []v1.Toleration{
			{
				Key:   "gpu",
				Value: "true",
			},
		},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		jobIds := make([]uuid.UUID, 0)
		for i := 0; i < numSmallCpuJobsToSchedule; i++ {
			jobId := uuid.New()
			jobIds = append(jobIds, jobId)
			report, err := db.SelectAndBindNodeToPod(jobId, smallCpuJob)
			if !assert.NoError(b, err) {
				return
			}
			if !assert.NotNil(b, report.Node) {
				return
			}
		}
		for i := 0; i < numLargeCpuJobsToSchedule; i++ {
			jobId := uuid.New()
			jobIds = append(jobIds, jobId)
			report, err := db.SelectAndBindNodeToPod(jobId, largeCpuJob)
			if !assert.NoError(b, err) {
				return
			}
			if !assert.NotNil(b, report.Node) {
				return
			}
		}
		for i := 0; i < numGpuJobsToSchedule; i++ {
			jobId := uuid.New()
			jobIds = append(jobIds, jobId)
			report, err := db.SelectAndBindNodeToPod(jobId, gpuJob)
			if !assert.NoError(b, err) {
				return
			}
			if !assert.NotNil(b, report.Node) {
				return
			}
		}

		// Release resources for the next iteration.
		for _, jobId := range jobIds {
			db.MarkJobRunning(jobId)
		}
	}
}

func BenchmarkSelectAndBindNodeToPod100(b *testing.B) {
	benchmarkSelectAndBindNodeToPod(70, 20, 10, 7, 2, 1, b)
}

func BenchmarkSelectAndBindNodeToPod1000(b *testing.B) {
	benchmarkSelectAndBindNodeToPod(700, 200, 100, 70, 20, 10, b)
}

func BenchmarkSelectAndBindNodeToPod10000(b *testing.B) {
	benchmarkSelectAndBindNodeToPod(7000, 2000, 1000, 700, 200, 100, b)
}

func TestAvailableByPriorityAndResourceType(t *testing.T) {
	tests := map[string]struct {
		Priorities     []int32
		UsedAtPriority int32
		Resources      map[string]resource.Quantity
	}{
		"lowest priority": {
			Priorities:     []int32{1, 5, 10},
			UsedAtPriority: 1,
			Resources: map[string]resource.Quantity{
				"cpu": resource.MustParse("1"),
				"gpu": resource.MustParse("2"),
			},
		},
		"mid priority": {
			Priorities:     []int32{1, 5, 10},
			UsedAtPriority: 5,
			Resources: map[string]resource.Quantity{
				"cpu": resource.MustParse("1"),
				"gpu": resource.MustParse("2"),
			},
		},
		"highest priority": {
			Priorities:     []int32{1, 5, 10},
			UsedAtPriority: 10,
			Resources: map[string]resource.Quantity{
				"cpu": resource.MustParse("1"),
				"gpu": resource.MustParse("2"),
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := NewAvailableByPriorityAndResourceType(tc.Priorities, tc.Resources)
			assert.Equal(t, len(tc.Priorities), len(m))

			m.MarkUsed(tc.UsedAtPriority, tc.Resources)
			for resourceType, quantity := range tc.Resources {
				for _, p := range tc.Priorities {
					actual := m.Get(p, resourceType)
					if p > tc.UsedAtPriority {
						assert.Equal(t, 0, quantity.Cmp(actual))
					} else {
						expected := resource.MustParse("0")
						assert.Equal(t, 0, expected.Cmp(actual))
					}
				}
			}

			m.MarkAvailable(tc.UsedAtPriority, tc.Resources)
			for resourceType, quantity := range tc.Resources {
				for _, p := range tc.Priorities {
					actual := m.Get(p, resourceType)
					assert.Equal(t, 0, quantity.Cmp(actual))
				}
			}
		})
	}
}

func TestAssignedByPriorityAndResourceType(t *testing.T) {
	tests := map[string]struct {
		Priorities     []int32
		UsedAtPriority int32
		Resources      map[string]resource.Quantity
	}{
		"lowest priority": {
			Priorities:     []int32{1, 5, 10},
			UsedAtPriority: 1,
			Resources: map[string]resource.Quantity{
				"cpu": resource.MustParse("1"),
				"gpu": resource.MustParse("2"),
			},
		},
		"mid priority": {
			Priorities:     []int32{1, 5, 10},
			UsedAtPriority: 5,
			Resources: map[string]resource.Quantity{
				"cpu": resource.MustParse("1"),
				"gpu": resource.MustParse("2"),
			},
		},
		"highest priority": {
			Priorities:     []int32{1, 5, 10},
			UsedAtPriority: 10,
			Resources: map[string]resource.Quantity{
				"cpu": resource.MustParse("1"),
				"gpu": resource.MustParse("2"),
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := NewAssignedByPriorityAndResourceType(tc.Priorities)
			assert.Equal(t, len(tc.Priorities), len(m))

			m.MarkUsed(tc.UsedAtPriority, tc.Resources)
			for resourceType, quantity := range tc.Resources {
				for _, p := range tc.Priorities {
					actual := m.Get(p, resourceType)
					if p <= tc.UsedAtPriority {
						assert.Equal(t, 0, quantity.Cmp(actual))
					} else {
						expected := resource.MustParse("0")
						assert.Equal(t, 0, expected.Cmp(actual))
					}
				}
			}

			m.MarkAvailable(tc.UsedAtPriority, tc.Resources)
			for resourceType := range tc.Resources {
				for _, p := range tc.Priorities {
					actual := m.Get(p, resourceType)
					expected := resource.MustParse("0")
					assert.Equal(t, 0, expected.Cmp(actual))
				}
			}
		})
	}
}
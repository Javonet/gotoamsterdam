package handler

import (
	"github.com/google/uuid"
	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/commandtype"
)

type ReferencesCache struct {
	referencesCacheDict map[string]interface{}
}

func NewReferenceCache() *ReferencesCache {
	return &ReferencesCache{referencesCacheDict: make(map[string]interface{})}
}

func (rc *ReferencesCache) GetInstance() *ReferencesCache {
	return rc
}

func (rc *ReferencesCache) CacheReference(reference interface{}) string {
	uuid := uuid.New()
	rc.referencesCacheDict[uuid.String()] = reference
	return uuid.String()
}

func (rc *ReferencesCache) ResolveReference(uuidString string, cmd *command.Command) (element interface{}, exists bool) {
	if cmd.CommandType == commandtype.Reference {
		element, exists = rc.referencesCacheDict[uuidString]
	} else {
		element, exists = nil, false
	}
	return
}

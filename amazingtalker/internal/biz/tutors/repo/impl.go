package repo

import (
	"time"

	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/blackhorseya/amazingtalker/pb"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	rw *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(rw *sqlx.DB) IRepo {
	return &impl{rw: rw}
}

func (i *impl) GetInfoBySlug(ctx contextx.Contextx, slug string) (*pb.Tutor, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tutor := pb.Tutor{}
	cmd := `select id, slug, name, headline, introduction from tutors where slug = ?`
	err := i.rw.GetContext(timeout, &tutor, cmd, slug)
	if err != nil {
		ctx.WithError(err).Error("get tutor by slug is failure")
		return nil, err
	}

	price := pb.PriceInfo{}
	cmd = `select trial_price as trial, normal_price as normal from tutor_lesson_prices where tutor_id = ?`
	err = i.rw.GetContext(timeout, &price, cmd, tutor.ID)
	if err != nil {
		ctx.WithError(err).Error("get price by tutor id is failure")
		return nil, err
	}
	tutor.PriceInfo = &price

	var languages []int32
	cmd = `select languages.id from languages
inner join tutor_languages tl on languages.id = tl.language_id
inner join tutors t on tl.tutor_id = t.id
where tl.tutor_id = ?
order by languages.id`
	err = i.rw.SelectContext(timeout, &languages, cmd, tutor.ID)
	if err != nil {
		ctx.WithError(err).Error("get languages by tutor id is failure")
		return nil, err
	}
	tutor.TeachingLanguages = languages

	return &tutor, nil
}

func (i *impl) ListByLangSlug(ctx contextx.Contextx, slug string) ([]*pb.Tutor, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var tutors []*pb.Tutor
	cmd := `select tutors.id, tutors.slug, name, headline, introduction from tutors
inner join tutor_languages tl on tutors.id = tl.tutor_id
inner join languages l on tl.language_id = l.id
where l.slug = ?`
	err := i.rw.SelectContext(timeout, &tutors, cmd, slug)
	if err != nil {
		ctx.WithError(err).Error("list tutors is failure")
		return nil, err
	}

	for _, tutor := range tutors {
		price := pb.PriceInfo{}
		cmd = `select trial_price as trial, normal_price as normal from tutor_lesson_prices where tutor_id = ?`
		err = i.rw.GetContext(timeout, &price, cmd, tutor.ID)
		if err != nil {
			ctx.WithError(err).Error("get price by tutor id is failure")
			return nil, err
		}
		tutor.PriceInfo = &price


		var languages []int32
		cmd = `select languages.id from languages
inner join tutor_languages tl on languages.id = tl.language_id
inner join tutors t on tl.tutor_id = t.id
where tl.tutor_id = ?
order by languages.id`
		err = i.rw.SelectContext(timeout, &languages, cmd, tutor.ID)
		if err != nil {
			ctx.WithError(err).Error("get languages by tutor id is failure")
			return nil, err
		}
		tutor.TeachingLanguages = languages
	}

	return tutors, nil
}
